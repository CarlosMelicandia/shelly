import { useQuery } from "@tanstack/react-query";
import { queryClient } from "@store/queryClient";

async function fetchUser() {
  const response = await fetch("http://localhost:8000/api/getUser", {
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch user session");
  }

  return response.json();
}

export default function useCurrentUser() {
  return useQuery(
    {
      queryKey: ["currentUser"],
      queryFn: fetchUser,
      retry: false,
      staleTime: 1000 * 60 * 5,
      onError: (error) => console.error("Error fetching user:", error),
    },
    queryClient,
  );
}
