import { useEffect, useState } from "react";
import getData from "../services/getData";
import putData from "../services/putData";

export default function useUser() {
  const [user, setUser] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchUser() {
    const _user = await getData("/user");
    setUser(_user.user);
  }

  async function updateUser(data) {
    setLoading(true);
    await putData("/user", data);
    await fetchUser();
    setLoading(false);
  }

  useEffect(() => {
    fetchUser();
  }, []);

  return [user, updateUser, loading];
}
