import { useApolloClient } from "@apollo/client";
import { useProjectQuery } from "../api/types.d";

import type { ProjectFile, Forbidden, ServerError } from "../api/types.d";

export function useProject(id: number) {
    const client = useApolloClient(); 
    return useProjectQuery({client, errorPolicy: "all", variables: { id }});
}

export default useProject;

export type ProjectFilesList = (
    { __typename: 'ProjectFilesList' }
    & { items: Array<(
        { __typename?: 'ProjectFile' }
        & Pick<ProjectFile, 'url'>
    )> }
    ) | (
    { __typename: 'Forbidden' }
    & Pick<Forbidden, 'message'>
    ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
)