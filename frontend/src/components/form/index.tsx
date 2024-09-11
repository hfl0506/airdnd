import { ComponentProps } from "react";
import { cn } from "../../utils/cn";

interface Props extends ComponentProps<"form"> {}

function Form({ children, className, ...rest }: Props) {
  return (
    <form
      className={cn(
        "w-[500px] min-h-[300px] flex flex-col gap-4 shadow-md p-8 rounded-md",
        className
      )}
      {...rest}
    >
      {children}
    </form>
  );
}

export default Form;
