import React from "react"
import { Anchor, AnchorExtendedProps } from "grommet"
import { Link, LinkProps } from "react-router-dom"

export type AnchorLinkProps = LinkProps & AnchorExtendedProps

const AnchorLink: React.FC<AnchorLinkProps> = (props) => {
    return (
        <Anchor as={Link} {...props} />
    )
}

export default AnchorLink