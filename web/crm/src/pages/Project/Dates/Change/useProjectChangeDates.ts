import { FetchResult, useApolloClient } from "@apollo/client"
import {
    useChangeProjectDatesMutation,
    ChangeProjectDatesMutation,
    ChangeProjectDatesInput,
    ChangeProjectDatesMutationResult,
} from "api/graphql"

export function useChangeProjectDates(): [
    (projectId: string, dates: ChangeProjectDatesInput) => Promise<FetchResult<ChangeProjectDatesMutation>>,
    ChangeProjectDatesMutationResult,
] {
    const client = useApolloClient()
    const [change, result] = useChangeProjectDatesMutation({ client, errorPolicy: "all" })

    return [(projectId: string, dates: ChangeProjectDatesInput) => change({ variables: { projectId, dates } }), result]
}
