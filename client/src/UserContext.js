import { createContext } from "react";

export const UserContext = createContext(localStorage.getItem("currentUser"))