import express from 'express'
import nodemailer from 'nodemailer'
import { SubmitFeedbackUseCase } from './use-cases/submit-feedback-use-case';
import { PrismaFeedbacksRepository } from './respositories/prisma/prisma-feedbacks-repository';

export const routes = express.Router()

const transport = nodemailer.createTransport({
  host: "smtp.mailtrap.io",
  port: 2525,
  auth: {
    user: "59e7eb94c2e377",
    pass: "bdb695da34efd6"
  }
});

routes.post('/feedbacks', async (req, res) => {
  const { type, comment, screenshot } = req.body

  const prismaFeedbacksRepository = new PrismaFeedbacksRepository
  const submitFeedbacUseCase = new SubmitFeedbackUseCase(prismaFeedbacksRepository)
  
  await submitFeedbacUseCase.execute({
    type,
    comment,
    screenshot
  })
  
  // await transport.sendMail({
  //   from: 'Equipe Team <sample@team.com>',
  //   to: 'Célio Viana <cl.juniorr@gmail.com>',
  //   subject: 'Novo feedback',
  //   html: [
  //     `<div style="font-family: sans-serif; font-size: 16px; color: #111">`,
  //     `<p>Tipo do feedback: ${type}</p>`,
  //     `<p>Comentário: ${comment}</p>`,
  //     `</div>`,
  //   ].join('\n')
  // })

  return res.status(201).send()
})