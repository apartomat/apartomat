import { FetchResult, useApolloClient } from "@apollo/client";
import { useCreateProjectMutation, CreateProjectMutation, CreateProjectMutationResult } from "../api/types.d";

export type CreateProjectFn = ({workspaceId, title}: {workspaceId: number, title: string }) => Promise<FetchResult<CreateProjectMutation>>;

export function useCreateProject(): [
    CreateProjectFn,
    CreateProjectMutationResult
] {
    const client = useApolloClient(); 
    const [create, result] = useCreateProjectMutation({ client, errorPolicy: 'all' });

    return [({workspaceId, title}) => create({ variables: { input: { workspaceId, title }} }), result];
}

export default useCreateProject;