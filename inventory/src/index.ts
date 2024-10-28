import { Page } from "playwright";
import { z } from "zod";
import { i } from "./inferable";
import * as inventory from "./inventory";

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
  name: "submitReservation",
  description: "Complete the reservation process for an item",
  schema: {
    input: z.object({
      customerName: z.string(),
      email: z.string(),
      quantity: z.number(),
      itemId: z.string(),
    }),
  },
  func: inventory.submitReservation,
});

service.register({
  name: "reserveItem",
  description: "Complete the reservation process for an item",
  schema: {
    input: z.object({
      customerName: z.string(),
      email: z.string(),
      quantity: z.number(),
      itemId: z.string(),
    }),
  },
  func: inventory.submitReservation,
});

service.start().then(() => {
  console.log("Inventory service started");
});
