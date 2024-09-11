export function nameCompose(name: string) {
  return name
    .split(" ")
    .map((n) => n.charAt(0).toUpperCase())
    .join("");
}
