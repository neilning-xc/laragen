<?php
namespace Tests\Feature;

use App\User;
use Dingo\Api\Routing\UrlGenerator;
use Tests\TestCase;

class UserApiTest extends TestCase
{
    
    public function setUp()
    {
        parent::setUp();
        factory(User::class)->create([
            'email' => 'test@test.com',
            'password' => bcrypt('123456')
        ]);
    }
    
    public function testLoginApi()
    {
        $url = app(UrlGenerator::class)->version('v1')->route('api.login');
        $response = $this->json('POST', $url, ['email' => 'test@test.com', 'password' => '123456']);
        $jsonData = $response->assertStatus(200)->assertSee('token')->getContent();
        $array = json_decode($jsonData, true);
        return $array['token'];
    }
    
    public function testGetUsersWithoutToken()
    {
        $url = app(UrlGenerator::class)->version('v1')->route('api.user.list');
        $response = $this->get($url);
        $response->assertStatus(401);
    }
    
    /**
     * @param $token
     * @depends testLoginApi
     */
    public function testGetUserWithToken($token)
    {
        $url = app(UrlGenerator::class)->version('v1')->route('api.user.list');
        $response = $this->withHeaders(['Authorization' => 'Bearer ' . $token])->json('GET', $url);
        $response->assertStatus(200);
    }
}
