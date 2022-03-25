import { useApolloClient } from "@apollo/client";
import { useSpecScreenQuery } from "../api/types.d";

export default function useSpecScreen(projectId: string) {
    const client = useApolloClient(); 
    return useSpecScreenQuery({client, errorPolicy: "all", variables: { projectId }});
}
