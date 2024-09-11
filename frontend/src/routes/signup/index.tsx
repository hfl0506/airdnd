import { createFileRoute } from "@tanstack/react-router";
import FormWrapper from "../../components/form/wrapper";
import Form from "../../components/form";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { RegisterInput, registerSchema } from "../../zod/auth";
import { useAuth } from "../../context/AuthContext";
import { useMutation } from "@tanstack/react-query";
import { registerApi } from "../../api/auth";
import { setTokens } from "../../utils/token";
import TextInput from "../../components/input";
import Typography from "../../components/typography";
import Button from "../../components/button";

export const Route = createFileRoute("/signup/")({
  component: SignUp,
});

function SignUp() {
  const navigate = Route.useNavigate();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<RegisterInput>({
    defaultValues: {
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
    resolver: zodResolver(registerSchema),
  });
  const { handleLogin } = useAuth();

  const registerMutate = useMutation({
    mutationKey: ["register"],
    mutationFn: async (data: RegisterInput) => await registerApi(data),
    onSuccess: (res) => {
      const { accessToken, refreshToken } = res.data.data;
      if (accessToken && refreshToken) {
        setTokens(accessToken, refreshToken);
        handleLogin();
      }
      reset();
      navigate({ to: "/", search: { page: 1, limit: 10, tab_id: "" } });
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const onSubmit = handleSubmit((data) => {
    registerMutate.mutate(data);
  });
  return (
    <FormWrapper>
      <Form onSubmit={onSubmit}>
        <TextInput label="Name:" type="text" {...register("name")} />
        {errors.name?.message && <Typography>{errors.name.message}</Typography>}
        <TextInput label="Email:" type="text" {...register("email")} />
        {errors.email?.message && (
          <Typography>{errors.email.message}</Typography>
        )}
        <TextInput
          label="Password:"
          type="password"
          {...register("password")}
        />
        {errors.password?.message && (
          <Typography>{errors.password.message}</Typography>
        )}
        <TextInput
          label="Confirm Password:"
          type="password"
          {...register("confirmPassword")}
        />
        {errors.confirmPassword?.message && (
          <Typography>{errors.confirmPassword.message}</Typography>
        )}
        <Button type="submit" className="mt-10">
          Sign Up
        </Button>
      </Form>
    </FormWrapper>
  );
}
