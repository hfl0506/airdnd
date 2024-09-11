import { ComponentProps } from "react";
import { cn } from "../../utils/cn";

interface Props extends ComponentProps<"button"> {}

function Button({ children, className, ...rest }: Props) {
  return (
    <button
      className={cn(
        className,
        "w-full rounded-md px-4 py-2 bg-red-400 text-white hover:bg-red-500"
      )}
      {...rest}
    >
      {children}
    </button>
  );
}

export default Button;
