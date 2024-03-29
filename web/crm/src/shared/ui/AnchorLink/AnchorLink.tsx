import { Anchor, AnchorProps } from "grommet"
import { Link, LinkProps } from "react-router-dom"

export type AnchorLinkProps = LinkProps & AnchorProps & Omit<JSX.IntrinsicElements["a"], "color">

export const AnchorLink: React.FC<AnchorLinkProps> = (props) => {
    return (
        <Anchor
            as={({ ...rest }) => {
                return <Link {...rest} />
            }}
            {...props}
        />
    )
}
