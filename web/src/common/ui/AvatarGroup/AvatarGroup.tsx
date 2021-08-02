import React from "react";

function AvatarGroup({ children }: { children: React.ReactChild[] }) {
    return (
        <div>
            {children.map((child) => {
                if (React.isValidElement(child)) {
                    return React.cloneElement(child, { style: { marginLeft: -8 } });
                }

                return child;
            })}
        </div>
    );
}

export default AvatarGroup;