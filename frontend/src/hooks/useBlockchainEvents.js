import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useBlockchainEvents() {
  const [blockchainEvents, setBlockchainEvents] = useState([]);
  async function fetchBlockchainEvents() {
    const _blockchainEvents = await getData("/blockchainEvents");
    setBlockchainEvents(_blockchainEvents.blockchainEvents);
  }
  useEffect(() => {
    fetchBlockchainEvents();
  }, []);
  return [blockchainEvents];
}
