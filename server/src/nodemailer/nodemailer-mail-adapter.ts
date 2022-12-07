import nodemailer from "nodemailer"
import { MailAdapter, SendMailData } from "../adapters/mail-adapter";

const transport = nodemailer.createTransport({
  host: "smtp.mailtrap.io",
  port: 2525,
  auth: {
    user: "59e7eb94c2e377",
    pass: "bdb695da34efd6"
  }
});

export class NodemailerMailAdapter implements MailAdapter {
  async sendMail({ subject, body }: SendMailData) {
    await transport.sendMail({
      from: 'Equipe Team <sample@team.com>',
      to: 'Célio Viana <cl.juniorr@gmail.com>',
      subject,
      html: body
    })
  }
}