import { useEffect, useState } from "react";

type MemoProps = {
  bookId: string;
  userId: string;
  onClose: () => void;
};

type MemoData = {
  id: number;
  userId: string;
  bookId: string;
  text: string;
  imageUrl?: string;
  createdAt: string;
  updatedAt: string;
};

// メモモーダル
const Memo = ({ bookId, userId, onClose }: MemoProps) => {
  const [memoText, setMemoText] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [imgFile, setImgFile] = useState<File | null>(null);

  useEffect(() => {
    const fetchMemo = async () => {
      setLoading(true);
      setError(null);

      try {
        const res = await fetch(
          `http://localhost:8080/api/memo?bookId=${bookId}&userId=${userId}`,
        );
        if (res.ok) {
          const data: { memo: MemoData | null } = await res.json();
          console.log("Fetched memo data:", data.memo?.text);
          setMemoText(data.memo?.text || "");
        } else if (res.status === 404) {
          setMemoText("");
        } else {
          setError("メモの取得に失敗しました");
        }
      } catch (err) {
        setError("予期せぬエラーが発生しました");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };
    fetchMemo();
  }, [bookId, userId]);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setImgFile(e.target.files[0]);
    } else {
      setImgFile(null);
    }
  };

  const handleSave = async () => {
    setLoading(true);
    setError(null);
    try {
      const formData = new FormData();
      formData.append("userId", userId);
      formData.append("bookId", bookId);
      formData.append("text", memoText);
      if (imgFile) {
        formData.append("image", imgFile);
      }

      const res = await fetch("http://localhost:8080/api/memo", {
        method: "POST",
        body: formData,
      });
      if (!res.ok) throw new Error();
      onClose();
    } catch {
      setError("Failed to save memo");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-items-center bg-black/10"
      role="dialog"
      aria-modal="true"
    >
      <div className="mx-auto w-80 rounded bg-white p-6 shadow-lg">
        <h4 className="test-lg mb-2 font-bold">Memo</h4>
        {loading ? (
          <p>Loading...</p>
        ) : (
          <>
            <textarea
              className="w-full rounded border p-2"
              rows={5}
              value={memoText}
              onChange={(e) => setMemoText(e.target.value)}
              placeholder="Let's Input memo!"
              required
            />
            <input
              type="file"
              accept="image/*"
              className="mt-2 w-full rounded border p-2"
              onChange={handleFileChange}
            />
            {error && <p className="mt-2 px-3 py-1 text-red-300">{error}</p>}
            <div className="mt-2 flex justify-end gap-2">
              <button
                className="rounded bg-gray-300 px-3 py-1"
                onClick={onClose}
              >
                Close
              </button>
              <button
                className="rounded bg-blue-500 px-3 py-1 text-white hover:bg-blue-600"
                onClick={handleSave}
                disabled={loading}
              >
                Save
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default Memo;
