import z from "zod";

export const listingSearchSchema = z.object({
  page: z.number().default(1),
  limit: z.number().default(10),
  tab_id: z.string().default(""),
});

export type ListingSearchInput = z.infer<typeof listingSearchSchema>;
