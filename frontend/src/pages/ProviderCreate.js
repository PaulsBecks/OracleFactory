import { useState } from "react";
import { Button } from "semantic-ui-react";
import ProviderCreateComponent from "../components/ProviderCreate";

export function ProviderCreate() {
  const [provider, setProvider] = useState({});
  return (
    <div>
      <ProviderCreateComponent provider={provider} setProvider={setProvider} />
    </div>
  );
}
