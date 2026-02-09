import { Box } from "grommet"

export function InnerHtmlPage({ scale, html }: { scale: number; html: string }) {
    return (
        <Box
            style={{
                transform: `scale(${scale})`,
                transformOrigin: "left top",
            }}
        >
            <div dangerouslySetInnerHTML={{ __html: html }} />
        </Box>
    )
}
