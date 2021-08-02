import React from "react";

const styles = {
    borderRadius: '50%',
    height: 32
};

function Avatar({ src, alt, className }: { src: string, alt?: string, className?: string }) {
    return (
        <img style={styles} src={src} alt={alt} className={className} />
    );
}

export default Avatar;