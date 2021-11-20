import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useBlockchainEvent(id) {
  const [blockchainEvent, setBlockchainEvent] = useState();

  async function fetchBlockchainEvent() {
    const data = await getData("/blockchainEvents/" + id);
    const _blockchainEvent = data.blockchainEvent;
    setBlockchainEvent(_blockchainEvent);
  }

  useEffect(() => {
    fetchBlockchainEvent();
  }, []); // eslint-disable-line

  return [blockchainEvent];
}
