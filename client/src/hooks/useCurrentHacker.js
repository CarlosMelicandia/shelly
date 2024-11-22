import { useQuery } from "@tanstack/react-query";
import { queryClient } from "@store/queryClient";

async function fetchHacker() {
  const response = await fetch("http://localhost:8000/api/getHacker", {
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch hacker application");
  }

  return response.json();
}

export default function useCurrentHacker() {
  return useQuery(
    {
      queryKey: ["currentHacker"],
      queryFn: fetchHacker,
      retry: false,
      staleTime: 1000 * 60 * 5,
      onError: (error) => console.error("Error fetching hacker application:", error),
    },
    queryClient,
  );
}

