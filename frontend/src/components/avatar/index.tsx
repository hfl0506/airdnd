import { nameCompose } from "../../utils/name";

type Props = {
  name: string;
  src?: string;
};

function Avatar({ name, src }: Props) {
  return (
    <>
      {src ? (
        <img src={src} className="w-8 h-8 rounded-full" alt={`${name} photo`} />
      ) : (
        <div className="flex justify-center items-center w-8 h-8 rounded-full bg-slate-500 text-white">
          {nameCompose(name)}
        </div>
      )}
    </>
  );
}

export default Avatar;
