import dayjs from "dayjs"

export function parseDatetimeFormat(date: Date) {
  return dayjs(date).format("YYYY-MM-DD");
}
