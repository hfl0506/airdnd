import {
  createContext,
  PropsWithChildren,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { getToken, removeTokens } from "../utils/token";
import { getMeApi, User } from "../api/auth";

type Context = {
  isAuth: boolean;
  handleLogin: () => void;
  handleLogout: () => void;
  user: User | null;
};

const authContext = createContext<Context | null>(null);

export const useAuth = () => {
  const context = useContext(authContext);
  if (!context) {
    throw new Error("auth context not found");
  }
  return context;
};

export function AuthProvider({ children }: PropsWithChildren) {
  const [isAuth, setIsAuth] = useState(!!getToken("access_token"));
  const [user, setUser] = useState<User | null>(null);

  const handleLogin = () => {
    setIsAuth(true);
  };

  const handleLogout = () => {
    removeTokens();
    setIsAuth(false);
  };

  const value = useMemo(() => {
    return {
      isAuth,
      user,
      handleLogin,
      handleLogout,
    };
  }, [isAuth, user, handleLogin, handleLogout]);

  const getMe = useCallback(async () => {
    try {
      const res = await getMeApi();
      if (res.data.success) {
        setUser(res.data.data);
      }
    } catch (error) {
      return null;
    }
  }, []);

  useEffect(() => {
    if (isAuth) {
      getMe();
    }
  }, [isAuth]);

  return <authContext.Provider value={value}>{children}</authContext.Provider>;
}
