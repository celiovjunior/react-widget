import {SubmitFeedbackUseCase} from '../use-cases/submit-feedback-use-case'

describe('Submit feedback', () => {
  it('should be able to submit a feedback', async () => {
    const submitFeedback = new SubmitFeedbackUseCase(
      { create: async () => {} },
      { sendMail: async () => {} }
    )

    await expect(submitFeedback.execute({
      type: 'BUG',
      comment: 'example comment',
      screenshot: 'data:image/png;base64fh4gs6d87h48df41hg56sdf4g'
    })).resolves.not.toThrow()
  })
})