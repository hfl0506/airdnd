import { zodResolver } from "@hookform/resolvers/zod";
import { createFileRoute } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import TextInput from "../../components/input";
import Button from "../../components/button";
import Typography from "../../components/typography";
import { useMutation } from "@tanstack/react-query";
import { LoginInput, loginSchema } from "../../zod/auth";
import { loginApi } from "../../api/auth";
import { setTokens } from "../../utils/token";
import { useAuth } from "../../context/AuthContext";
import FormWrapper from "../../components/form/wrapper";
import Form from "../../components/form";

export const Route = createFileRoute("/login/")({
  component: Login,
});

function Login() {
  const navigate = Route.useNavigate();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<LoginInput>({
    defaultValues: {
      email: "",
      password: "",
    },
    resolver: zodResolver(loginSchema),
  });
  const { handleLogin } = useAuth();

  const loginMutate = useMutation({
    mutationKey: ["login"],
    mutationFn: async (data: LoginInput) => await loginApi(data),
    onSuccess: (res) => {
      const { accessToken, refreshToken } = res.data.data;
      if (accessToken && refreshToken) {
        setTokens(accessToken, refreshToken);
        handleLogin();
      }
      reset();
      navigate({ to: "/", search: { page: 1, limit: 1, tab_id: "" } });
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const onSubmit = handleSubmit((data) => {
    loginMutate.mutate(data);
  });

  return (
    <FormWrapper>
      <Form onSubmit={onSubmit}>
        <TextInput label="Email:" type="text" {...register("email")} />
        {errors.email?.message && (
          <Typography>{errors.email.message}</Typography>
        )}
        <TextInput label="Password" type="password" {...register("password")} />
        {errors.password?.message && (
          <Typography>{errors.password.message}</Typography>
        )}
        <Button type="submit" className="mt-10">
          Login
        </Button>
      </Form>
    </FormWrapper>
  );
}
