import { join } from "path";
import { Browser, By, Builder } from "selenium-webdriver"
import { writeFile, mkdir } from "fs/promises";

const URL = 'https://www.rocketleague.com/en/developer/stats-api';
const JSON_DEST = join(import.meta.dirname, '..', 'data');

async function openPage(url) {
    const driver = await new Builder().forBrowser(Browser.CHROME).build();
    await driver.get(url);
    return driver;
}


async function extractEventNames(driver) {
    const events = await driver.findElements(By.css('h3 h3'));
    const eventNames = await Promise.all(events.map(e => e.getText()));
    return eventNames;
}

async function extractEventDescriptions(driver,) {
    const eventNames = await extractEventNames(driver);
    const eventDescriptions = await Promise.all(eventNames.map(async (name) => {
        const descriptionElement = await driver.findElement(By.css(`#${name} > div > div > span:first-child`)).getText();
        return descriptionElement;
    }))
    return eventDescriptions;
}

async function getFieldPadding(row) {
    const paddingStr = await row.findElement(By.css('td:first-child')).getCssValue('padding-inline-start')
    return parseInt(paddingStr.replace("px", ""));
}

async function getCellText(row, columnIndex) {
    const cell = await row.findElement(By.css(`td:nth-child(${columnIndex + 1})`));
    return await cell.getText();
}

async function parseEventField(rows, index) {
    // Parse field and all of its subfields
    // Return [field name, field structure, next index]
    const struct = {
        "fieldType": "",
        "description": "",
        "subFields": {},
    }

    const name = await getCellText(rows[index], 0);
    const typeStr = await getCellText(rows[index], 1);
    const description = await getCellText(rows[index], 2);

    // TODO: Parse conditional/spectator flag in description
    struct.fieldType = typeStr;
    struct.description = description;

    // Check for basic type (!= object or array)
    if (typeStr !== "object" && typeStr !== "array") {

        return [name, struct, index + 1]
    }

    // Need to parse subfields, so search for subsequent fields with > padding
    const padding = await getFieldPadding(rows[index]);
    index++;
    while (index < rows.length && await getFieldPadding(rows[index]) > padding) {
        let subName, subStruct;
        [subName, subStruct, index] = await parseEventField(rows, index, padding);
        struct.subFields[subName] = subStruct;
    }

    return [name, struct, index];
}

async function parseEventTable(table) {
    // Parse table and return event structure
    let rows = await table.findElements(By.css('tr'));
    rows = rows.slice(1);
    let idx = 0;

    // Process each row to extract the field name and type
    // Consider case where field is an object/array:
    //   All subfields will be indented, which must be checked with padding-inline-start property
    const out = {};
    while (idx < rows.length) {
        let struct;
        let name;
        [name, struct, idx] = await parseEventField(rows, idx, await getFieldPadding(rows[idx]));
        out[name] = struct;
    }
    return out;
}

async function extractEventStructures(driver) {
    /* Event structures returned as:
     * Dict of event names to their structures, which will be stored as:
     * {
     *   "fieldName1": {
     *     "fieldType": "fieldType1", (fieldType1* = optional type, fieldType1? = spectator type)
     *     "description": "Description of the field",
     *     "subFields": {
     *       ...
     *     }*
     *   },
     *   ...
     * }
    */

    const eventNames = await extractEventNames(driver);

    // Search for div containing event name and extract contents of <table>
    const eventStructures = {}
    for (const name of eventNames) {
        const table = await driver.findElement(By.css(`#${name} table`));
        eventStructures[name] = await parseEventTable(table);
    }

    return eventStructures;
}

async function main() {
    console.log('Opening Stats API page...');
    let driver;

    try {
        // Ensuring json directory exists
        console.log(`Ensuring ${JSON_DEST} exists...`);
        await mkdir(JSON_DEST, { recursive: true });

        driver = await openPage(URL);
        console.log(`Page ${URL} opened.`);

        console.log(`Extracting event names into ${JSON_DEST}/eventnames.json`);

        const eventNames = await extractEventNames(driver);
        await writeFile(`${JSON_DEST}/eventnames.json`, JSON.stringify(eventNames));

        console.log(`Event names successfully written to ${JSON_DEST}/eventnames.json`);

        console.log(`Extracting event descriptions into ${JSON_DEST}/eventdescriptions.json`);

        const eventDescriptions = await extractEventDescriptions(driver);
        await writeFile(`${JSON_DEST}/eventdescriptions.json`, JSON.stringify(eventDescriptions));

        console.log(`Event descriptions successfully written to ${JSON_DEST}/eventdescriptions.json`);

        console.log(`Extracting event structures into ${JSON_DEST}/eventstructures.json`);

        const eventStructures = await extractEventStructures(driver);
        await writeFile(`${JSON_DEST}/eventstructures.json`, JSON.stringify(eventStructures));

        console.log(`Event structures successfully written to ${JSON_DEST}/eventstructures.json`);

        console.log(`Closing ${URL}...`);
        await driver.quit();
        console.log(`Page ${URL} closed successfully.`);

    } catch (error) {
        console.error('Error:', error.message);
    } finally {
        try {
            await driver.quit();
        } catch (error) {
            return;
        }
    }
}

main();