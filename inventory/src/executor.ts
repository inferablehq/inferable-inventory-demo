import { i } from "./inferable";

async function reserveItems(
  itemName: string,
  quantity: number,
  email: string,
  address: string
): Promise<{
  orderId: string;
  customerId: string;
  reservationId: string;
} | null> {
  const run = await i.run({
    message: `Create a new order for ${itemName} with ${quantity} items. Reserve the stock from the inventory before creating the order. The customer's email is ${email} and their address is ${address}.`,
    result: {
      schema: {
        type: "object",
        properties: {
          orderId: {
            type: "string",
          },
          customerId: {
            type: "string",
          },
          reservationId: {
            type: "string",
          },
        },
        required: ["orderId", "customerId", "reservationId"],
      },
    },
  });

  async function poll() {
    let result = await run.poll();
    if (result?.status === "done" || result?.status === "failed") {
      return result.result as {
        orderId: string;
        customerId: string;
        reservationId: string;
      } | null;
    }
    await new Promise((resolve) => setTimeout(resolve, 1000));
    return poll();
  }

  return poll();
}

reserveItems("Tablet", 1, "john.doe@example.com", "123 Main St, Anytown, USA")
  .then((result) => {
    console.log(result);
  })
  .catch((error) => {
    console.error(error);
  });
