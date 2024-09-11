import z from "zod"

export const createReviewSchema = z.object({
  rating: z.number(),
  roomId: z.string(),
  comment: z.string().min(1, "comment required").min(3, "comment at least 3 characters")
})

export type CreateReviewInput = z.infer<typeof createReviewSchema>
