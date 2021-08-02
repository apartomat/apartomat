import { FetchResult, useApolloClient } from "@apollo/client";
import { useUploadProjectFileMutation, UploadProjectFileMutation, UploadProjectFileMutationResult } from "../api/types.d";

export type UploadFn = ({projectId, file}: { projectId: number, file: File}) => Promise<FetchResult<UploadProjectFileMutation>>;

export function useUploadProjectFile(): [
    UploadFn,
    UploadProjectFileMutationResult
] {
    const client = useApolloClient(); 
    const [upload, result] = useUploadProjectFileMutation({ client, errorPolicy: 'all' });

    return [({projectId, file}) => upload({ variables: { file: { projectId, file } } }), result];
}

export default useUploadProjectFile;