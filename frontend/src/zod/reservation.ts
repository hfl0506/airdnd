import z from "zod";

export const reservationSchema = z.object({
  startDate: z.date(),
  endDate: z.date(),
  roomId: z.string(),
  price: z.number(),
  adult: z.string().default("0"),
  child: z.string().default("0"),
  infant: z.string().default("0")
}).refine(data => data.startDate < data.endDate, "Start date should earlier than End date")

export type ReservationSchema = z.infer<typeof reservationSchema>;

export type ReservationInput = {
  startDate: string;
  endDate: string;
  roomId: string;
  price: number;
  checkIn: {
    adult: number;
    child: number;
    infant: number;
  }
}
