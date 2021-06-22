import useUser from "../hooks/useUser";
import { UserForm } from "../components";
import { useEffect, useState } from "react";
import { Button } from "semantic-ui-react";

export default function Settings() {
  const [user, updateUser, loading] = useUser();
  const [localUser, setLocalUser] = useState();

  useEffect(() => {
    setLocalUser(user);
  }, [user]);

  console.log(localUser);

  return (
    <div>
      <h1>Settings</h1>
      <UserForm user={localUser} setUser={setLocalUser} />
      <br />
      {JSON.stringify(localUser) !== JSON.stringify(user) && (
        <>
          <br />
          <Button
            loading={loading}
            positive
            basic
            onClick={() => updateUser(localUser)}
            content="Save"
          />
        </>
      )}
    </div>
  );
}
