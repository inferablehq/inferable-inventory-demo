import { Inferable } from "inferable";
import { z } from "zod";
import { chromium, Page } from "playwright";
import * as inventory from "./inventory";

// Instantiate the Inferable client.
const i = new Inferable({
  // To get a new key, run:
  // npx @inferable/cli auth keys create 'My New Machine Key' --type='cluster_machine'
  apiSecret: process.env.INFERABLE_API_SECRET,
});

const service = i.service({
  name: "inventory",
});

let page: Page;

service.register({
  name: "login",
  description: "Log in to the inventory system",
  schema: {
    input: z.object({}),
  },
  func: inventory.login,
});

service.register({
  name: "getInventoryData",
  description: "Retrieve inventory data from the system",
  schema: {
    input: z.object({}),
  },
  func: inventory.getInventoryData,
});

service.register({
  name: "clickReserveButton",
  description: "Click the reserve button for a specific item",
  schema: {
    input: z.object({ itemId: z.string() }),
  },
  func: inventory.clickReserveButton,
});

service.register({
  name: "completeReservation",
  description: "Complete the reservation process for an item",
  schema: {
    input: z.object({
      customerName: z.string(),
      email: z.string(),
      quantity: z.number(),
    }),
  },
  func: inventory.completeReservation,
});

service.start().then(() => {
  console.log("Inventory service started");
});
