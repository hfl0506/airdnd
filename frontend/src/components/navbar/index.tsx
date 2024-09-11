import { Link } from "@tanstack/react-router";
import { useAuth } from "../../context/AuthContext";
import Avatar from "../avatar";

function Navbar() {
  const { isAuth, user, handleLogout } = useAuth();
  return (
    <nav className="w-full h-14 px-10 flex items-center justify-between">
      <Link
        className="text-red-300 text-3xl"
        to="/"
        search={{ page: 1, limit: 10, tab_id: "" }}
      >
        airdnd
      </Link>
      <ul className="flex gap-4 items-center">
        {isAuth ? (
          <>
            <Link to="/bookings">Bookings</Link>
            <Link to="/wishlist">Wishlist</Link>
            <Avatar name={user?.name ?? "N/A"} src={user?.photo!} />
            <button type="button" onClick={handleLogout}>
              Logout
            </button>
          </>
        ) : (
          <>
            <Link to="/login">Login</Link>
            <Link to="/signup">Sign Up</Link>
          </>
        )}
      </ul>
    </nav>
  );
}

export default Navbar;
