require("dotenv").config();

import MailosaurClient from "mailosaur";

if (!process.env.MAILOSAUR_API_KEY) {
  throw new Error("MAILOSAUR_API_KEY is not set");
}

if (!process.env.MAILOSAUR_SERVER_ID) {
  throw new Error("MAILOSAUR_SERVER_ID is not set");
}

if (!process.env.MAILOSAUR_EMAIL_ADDRESS) {
  throw new Error("MAILOSAUR_EMAIL_ADDRESS is not set");
}

const mailosaur = new MailosaurClient(process.env.MAILOSAUR_API_KEY);

const processedEmails = new Set<string>();

export async function pollForEmails(
  emailAddress: string = process.env.MAILOSAUR_EMAIL_ADDRESS!
) {
  const response = await mailosaur.messages
    .get(
      process.env.MAILOSAUR_SERVER_ID!,
      {
        sentTo: emailAddress,
      },
      {
        timeout: 20000, // 20 seconds (in milliseconds)
      }
    )
    .catch((error) => {
      console.error(error);
    });

  if (response) {
    // if received in the last 1 minute
    if (
      response.received &&
      new Date().getTime() - response.received.getTime() < 60000 &&
      response.id &&
      !processedEmails.has(response.id)
    ) {
      processedEmails.add(response.id);

      console.log(response);
    }
  }

  // wait for 5 seconds
  await new Promise((resolve) => setTimeout(resolve, 5000));

  return pollForEmails(emailAddress);
}

pollForEmails();
