"use client"
import { signIn } from "next-auth/react"
import { useState } from "react"

export default function SigninForm() {
  const [student_id, setStudent_id] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  

  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();

    // const result = await signIn("credentials", { student_id, password, redirect: false, callbackUrl: "/" });
    const result = await userSignin(student_id, password);
    // console.log("result",result);
    SetCookie("jwt", result?.data);

    if(result?.error) {
      setError("Invalid credentials");
    }else{
      console.log("Sign-in successful", result);
      document.cookie = `jwt=${result}; path=/; Secure; SameSite=Strict`;
    }

  };

    return(

        <div className="flex bg-white rounded-lg shadow-lg">
  
          {/* Form Section */}
          <form className="flex flex-col flex-1 p-6 space-y-4"
            onSubmit={handleSignIn}
          >
            <div>
                <div className="text-black font-poppins text-[14px] sm:text-[14px] lg:text-[16px]">StudentID</div>
                <input
                type="text"
                name="studentID"
                // placeholder="Student ID"
                onChange={(e) => setStudent_id(e.target.value)}
                className="w-full px-4 py-2 border bg-gray-300 text-[14px] sm:text-[14px] lg:text-[16px] text-black rounded-xl"
                required
                />
            </div>

            <div>
                <div className="text-black font-poppins text-[14px] sm:text-[14px] lg:text-[16px]">Password</div>
                <input
                type="password"
                name="password"
                // placeholder="Student ID"
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-4 py-2 border bg-gray-300 text-[14px] sm:text-[14px] lg:text-[16px] text-black rounded-xl"
                required
                />
            </div>
  
            
            {/* Submit Button */}
            <button
              type="submit"
              className="w-full bg-[#E799AC] text-white py-2 rounded-xl hover:bg-pink-600"
            >
              Sign in
            </button>
            {error && <p className="text-red-500">{error}</p>}
          </form>
        </div>

    )

}
