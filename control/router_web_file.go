package control

var router_head = `<?php
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\Facades\App;
Auth::routes();

Route::group(['middleware' => ['auth'], 'prefix' => 'admin', 'namespace' => "Admin"], function () {

    Route::get('/dashboard', 'DashboardController@dashboard')->name('dashboard');   #後台首頁
    Route::resource('roles', 'RoleController');                                 #用戶組/角色
    Route::resource('users', 'UserController');                                 #後台管理員
    Route::resource('posts', 'PostController');                                 #新聞管理
    Route::resource('postCategories', 'PostCategoryController');                #新闻分类管理
`

var router_footer = `
    Route::post('delimage', 'TemplateController@delimage')->name('delimage');               //刪除圖片
    Route::post('destroy_image', 'TemplateController@remove_image')->name('destroy_image'); //刪除圖片
    Route::get('google2faSet/{id}',  'UserController@google2faSet')->name('users.google2faSet');
});
Route::get('/', 'HomeController@index')->name('index');
Route::get('/home', 'HomeController@index')->name('home');
Route::post('uploadimgs', 'HomeController@uploadimgs')->name('uploadimgs');
Route::get('/setcn', function(){
    App::setLocale('cn');
    return redirect()->back();
});

Route::get('/seten', function(){
    App::setLocale('en');
    return redirect()->back();
});`
