import { FetchResult, useApolloClient } from "@apollo/client"
import { useChangeProjectDatesMutation, ChangeProjectDatesMutation, ChangeProjectDatesInput, ChangeProjectDatesMutationResult } from "../api/types.d"

export function useChangeDates(): [
    (projectId: string,data: ChangeProjectDatesInput) => Promise<FetchResult<ChangeProjectDatesMutation>>,
    ChangeProjectDatesMutationResult
] {
    const client = useApolloClient()
    const [change, result] = useChangeProjectDatesMutation({ client, errorPolicy: 'all' })

    return [
        (projectId: string, dates: ChangeProjectDatesInput) => change({ variables: { projectId, dates } }),
        result,
    ]
}
