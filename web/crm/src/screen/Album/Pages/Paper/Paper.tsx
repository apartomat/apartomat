import {Box, BoxExtendedProps} from "grommet";

export default function Paper({
   children,
   size = "A4",
   scale = 1.0,
   ...boxProps
}: {
    children: JSX.Element | never[] | undefined | string,
    size?: "A4" | "A5",
    scale?: 0.05 | 0.1 | 0.25 | 0.3 | 0.4 | 0.5 | 0.6 | 0.7 | 0.75 | 1
} & BoxExtendedProps) {
    return (
        <Box {...boxProps} width={`calc(${scale} * 210mm)`} height={`calc(${scale} * 297mm)`}>
            {children}
        </Box>
    )
}
