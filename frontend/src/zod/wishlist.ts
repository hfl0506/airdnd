import z from "zod"

export const upsertWishlistSchema = z.object({
  roomId: z.string(),
  liking: z.boolean()
})

export type UpsertWishlistInput = z.infer<typeof upsertWishlistSchema>;
