type props = {
    imageUrl: string;
};

export default function ({ imageUrl }: props) {
    return (
        <img
            className="w-40 h-60 rounded-l border-gray-700 border shadow-xl"
            src={imageUrl}
        />
    );
}
