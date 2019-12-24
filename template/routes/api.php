<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});

$api = app(\Dingo\Api\Routing\Router::class);

$api->version('v1', function ($api) {
    $api->post('/login', '\App\Http\Controllers\Api\LoginController@login')->name('api.login');
    
    $api->group(['middleware' => 'api.auth'], function ($api) {
        $api->get('/user', '\App\Http\Controllers\Api\UserController@index')->name('api.user.list');
        $api->post('/user', '\App\Http\Controllers\Api\UserController@store')->name('api.user.create');
        $api->put('/user/{id}', '\App\Http\Controllers\Api\UserController@update')->name('api.user.update');
        $api->delete('/user/{id}', '\App\Http\Controllers\Api\UserController@destroy')->name('api.user.delete');
    });
});
