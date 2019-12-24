<?php
/**
 * Time: 2019-11-27 22:42:05
 * Author: Neil Xuchu Ning
 * Email: neil.ning@chinanetcloud.com
 * Desc: LoginController.php
 */
namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;
use Tymon\JWTAuth\Exceptions\JWTException;
use Tymon\JWTAuth\Facades\JWTAuth;

class LoginController
{
    public function login(Request $request)
    {
        $credentials = $request->only('email', 'password');
        try {
            if (! $token = JWTAuth::attempt($credentials)) {
                throw new UnauthorizedHttpException("Credential error");
            }
        } catch (JWTException $exception) {
            throw new UnauthorizedHttpException("Token error");
        }
        return compact('token');
    }
}
