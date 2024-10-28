import { chromium, Page } from "playwright";

let page: Page;

export const login = async () => {
  const browser = await chromium.launch({
    headless: false,
    timeout: 10000,
    slowMo: 1000,
  });
  page = await browser.newPage();

  await page.goto("http://127.0.0.1:5556/");

  // Fill in the username and password
  await page.type("#username", "admin", { delay: 100 });
  await page.type("#password", "admin", { delay: 100 });

  // Click the login button
  await page.click('input[type="submit"]');

  // Check if login was successful
  const currentUrl = page.url();

  const success = currentUrl.includes("inventory.html");

  if (success) {
    console.log("Login successful");
  } else {
    console.log("Login failed");
  }

  return success;
};

export const getInventoryData = async (): Promise<InventoryItem[]> => {
  if (!page) {
    throw new Error("Page is not initialized. Please login first.");
  }

  // Wait for the inventory table to load
  await page.waitForSelector("#inventoryTable");

  // scroll to the bottom of the page
  await page.evaluate(() => {
    window.scrollTo(0, document.body.scrollHeight);
  });

  // pause 1s for demo
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Get all rows from the table except the header
  const rows = await page.$$("#inventoryTable tr:not(:first-child)");

  const inventoryData: InventoryItem[] = [];

  for (const row of rows) {
    const cells = await row.$$("td");
    const item: InventoryItem = {
      id: await cells[0].innerText(),
      name: await cells[1].innerText(),
      category: await cells[2].innerText(),
      brand: await cells[3].innerText(),
      model: await cells[4].innerText(),
      quantity: parseInt(await cells[5].innerText(), 10),
      unitPrice: parseFloat((await cells[6].innerText()).replace("$", "")),
      supplier: await cells[7].innerText(),
    };
    inventoryData.push(item);
  }

  return inventoryData;
};

export const submitReservation = async ({
  customerName,
  email,
  quantity,
  itemId,
}: {
  customerName: string;
  email: string;
  quantity: number;
  itemId: string;
}): Promise<{ reservationId: string }> => {
  if (!page) {
    throw new Error("Page is not initialized. Please login first.");
  }

  // Find and click the "Reserve" link for the specified item
  const reserveLink = await page.$(`a#${itemId}`);
  if (!reserveLink) {
    throw new Error(`Reserve link not found for item "${itemId}".`);
  }

  await reserveLink.click();
  console.log(`Clicked Reserve link for item: ${itemId}`);

  // Wait for the reservation form to appear
  await page.waitForSelector("#reservationForm");

  // Fill in the customer name
  await page.type("#reservationForm input[name='name']", customerName, {
    delay: 100,
  });

  // Fill in the email
  await page.type("#reservationForm input[name='email']", email, {
    delay: 100,
  });

  // Fill in the quantity
  await page.type(
    "#reservationForm input[name='quantity']",
    quantity.toString(),
    {
      delay: 100,
    }
  );

  // Click the "Reserve Now!" button
  await page.click("#reservationForm input[type='submit']");

  // Wait for the success message to appear
  await page.waitForSelector("#successMessage", { state: "visible" });

  const successMessage = await page.$("#successMessage h2");
  if (successMessage) {
    const messageText = await successMessage.innerText();
    console.log(`Reservation completed: ${messageText}`);
  } else {
    console.log(
      "Reservation completed, but couldn't retrieve the success message."
    );
  }

  return {
    reservationId: "123",
  };
};

interface InventoryItem {
  id: string;
  name: string;
  category: string;
  brand: string;
  model: string;
  quantity: number;
  unitPrice: number;
  supplier: string;
}

// Example usage:
// async function run() {
//   await login();
//   const inventoryData = await getInventoryData();
//   console.log(inventoryData);

//   // Example of reserving an item
//   const itemToReserve = { itemId: "O003" };
//   await clickReserveButton(itemToReserve);
//   await completeReservation({
//     customerName: "John Doe",
//     email: "john.doe@example.com",
//     quantity: 2,
//   });
// }

// run().catch(console.error);
