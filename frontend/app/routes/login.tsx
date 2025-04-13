import { Form } from "@remix-run/react";
import { useMutation } from "@tanstack/react-query";
import { postAuthGoogle } from "~/client/default/default";

export default function LoginPage() {
  const handleLogin = () => {
    window.location.href = "/auth/google";
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-4 bg-white rounded shadow-md">
        <h1 className="text-2xl font-bold text-center">Googleでログイン</h1>
        <button
          onClick={handleLogin}
          className="w-full px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300"
        >
          Googleでログイン
        </button>
      </div>
    </div>
  );
}
