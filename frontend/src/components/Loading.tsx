// import "../../scss/components/Loading.scss";FIXME:tailwind css

const Loading = () => {
  return (
    <div className="loading-circle" role="status" aria-live="polite">
      <div className="loading-circle-border">
        <div className="loading-circle-core"></div>
      </div>
    </div>
  );
};

export default Loading;