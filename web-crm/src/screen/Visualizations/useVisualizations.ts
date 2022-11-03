import { useApolloClient } from "@apollo/client"
import { ProjectVisualizationsListFilter, useVisualizationsScreenQuery } from "api/types.d"
import { useState } from "react";

export { VisualizationStatus } from  "api/types.d"

export function useVisualizations(id: string, filter: ProjectVisualizationsListFilter) {

    const [ first, setFirst] = useState(true)

    const client = useApolloClient();

    const result = useVisualizationsScreenQuery({
        client,
        errorPolicy: "all",
        variables: { id, filter },
        notifyOnNetworkStatusChange: true
    })

    const { called, loading } = result

    if (called && first && !loading) {
        setFirst(false)
    }

    return { ...result, first }
}