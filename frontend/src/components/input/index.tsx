import { ComponentProps, forwardRef } from "react";
import { cn } from "../../utils/cn";

interface Props extends ComponentProps<"input"> {
  label: string;
}

const TextInput = forwardRef<HTMLInputElement, Props>(function TextInput(
  { label, className, ...rest },
  ref
) {
  return (
    <label className="flex flex-col gap-2">
      {label}
      <input
        ref={ref}
        className={cn(
          className,
          "w-full rounded-md h-10 indent-2 border-[1px] border-slate-300 focus:outline-none focus:ring-2 focus:ring-red-300 focus:border-0"
        )}
        {...rest}
      />
    </label>
  );
});

TextInput.displayName = "TextInput";

export default TextInput;
