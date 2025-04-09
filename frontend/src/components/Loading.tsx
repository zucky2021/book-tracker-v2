const Loading = () => {
  return (
    <div
      className="mx-auto flex h-32 w-32 items-center justify-center bg-transparent sm:h-40 sm:w-40"
      role="status"
      aria-live="polite"
    >
      <div className="relative flex justify-center items-center w-20 h-20 p-1.5 rounded-full bg-gradient-to-b from-cyan-100 to-cyan-500 animate-spin">
        <div className="w-full h-full bg-white rounded-full"></div>
      </div>
    </div>
  );
};

export default Loading;
