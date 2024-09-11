import { ComponentProps } from "react";

interface Props extends ComponentProps<"span"> {}

function Typography({ children }: Props) {
  return <span className="text-red-500">{children}</span>;
}

export default Typography;
