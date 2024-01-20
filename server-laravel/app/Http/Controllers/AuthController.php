<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\DB;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Exception;
use Illuminate\Support\Facades\Auth;


class AuthController extends Controller
{
    public function __construct()
    {
        $this->middleware('auth:api', ['except' => ['login', 'register', 'getAllMahasiswa']]);
    }

    private function generateJWT($mahasiswa)
    {
        $expirationTime = time() + 3600*24*365; // jwt valid for 1 hour from the issued time
        $credentials = array(
            "id" =>              $mahasiswa->id,
            "password" =>        $mahasiswa->password,
            "username" =>        $mahasiswa->username,
            "email" =>           $mahasiswa->email,
            "nim" =>             $mahasiswa->nim,
            // "expiration_date" =>    $expirationTime,
        );

        if (! $token = auth()->attempt($credentials)) {
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        return $token;
    }

    public function login(Request $request)
    {
        $content = $request->getContent();
        $email = $request->input('email');

        // find email
        $mahasiswa = DB::selectOne("SELECT * FROM user_api_mahasiswa WHERE email = '$email'");
        if ($mahasiswa == null) {
            return response()->json([
                'message' => 'Email not found'
            ], 404);
        }

        // Check hashed password
        $password = $request->input('password');
        if (!password_verify($password, $mahasiswa->password)) {
            return response()->json([
                'message' => 'Wrong password'
            ], 401);
        }

        // Update last_login
        $lastLogin = date("Y-m-d H:i:s");
        try {
            DB::update("UPDATE user_api_mahasiswa SET last_login = '$lastLogin' WHERE email = '$email'");
        }
        catch (Exception $e) {
            $message = $e->getMessage();
            return response()->json(["message" => $message], 500);
        }

        // Generate JWT
        $mahasiswa->password = $password; // because Auth::attemp for some reason, hash it
        $accessToken = $this->generateJWT($mahasiswa);

        return response()->json([
            'id' => $mahasiswa->id, 
            'password' => $password, 
            'last_login' => $mahasiswa->last_login, 
            'is_superuser' => $mahasiswa->is_superuser, 
            'username' => $mahasiswa->username, 
            'first_name' => $mahasiswa->first_name, 
            'last_name' => $mahasiswa->last_name, 
            'email' => $email, 
            'is_staff' => $mahasiswa->is_staff, 
            'is_active' => $mahasiswa->is_active, 
            'date_joined' => $mahasiswa->date_joined, 
            'nim' => $mahasiswa->nim, 
            'access_token' => $accessToken,
        ]);
    }
    public function register(Request $request)
    {
        $content = $request->getContent();
        $email = $request->input('email');
        
        // check if email already exists
        $mahasiswa = DB::select("SELECT * FROM user_api_mahasiswa WHERE email = '$email'");
        if (count($mahasiswa) > 0) {
            return response()->json([
                'message' => 'Email already exists'
            ], 409);
        }

        // Assign values to newMahasiswa
        $dateJoined = date("Y-m-d H:i:s");
        $lastLogin = $dateJoined;
        $password = $request->input('password');
        $password = password_hash($password, PASSWORD_DEFAULT);
        $username = $request->input('username');
        $nim = $request->input('nim');

        // Insert newMahasiswa to database
        try {
            DB::insert("INSERT INTO user_api_mahasiswa (password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined, nim) VALUES ('$password', '$lastLogin', 'false', '$username', '', '', '$email', 'false', 'true', '$dateJoined', '$nim')");
        }
        catch (Exception $e) {
            $message = $e->getMessage();
            return response()->json(["message" => $message], 500);
        }

        return $content;
    }
    public function authenticate(Request $request)
    {
        $mahasiswa = Auth::user();
        return $mahasiswa;
    }
    public function getAllMahasiswa()
    {
        try{
            $mahasiswa = DB::select("SELECT * FROM user_api_mahasiswa");
        }
        catch (Exception $e) {
            return response()->json(["message" => $e->getMessage()], 500);
        }
        return $mahasiswa;
    }
}
