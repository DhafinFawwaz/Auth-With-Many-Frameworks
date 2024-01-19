<?php

use App\Http\Controllers\AuthController;
use Illuminate\Database\QueryException;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

// public routes
Route::post('/login', [AuthController::class, 'login']);

Route::post('/register', [AuthController::class, 'register']);

Route::get('/login', [AuthController::class, 'getAllMahasiswa']);

// protected routes
Route::post('/', [AuthController::class, 'authenticate']);