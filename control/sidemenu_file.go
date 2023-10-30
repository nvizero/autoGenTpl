package control

var sidemenu_head = `<div class="vertical-menu">
    <div data-simplebar class="h-100">
        <!--- Sidemenu -->
        <div id="sidebar-menu">
            <!-- Left Menu Start -->
            <ul class="metismenu list-unstyled" id="side-menu">
                <li class="menu-title">{{ __('menu.normal') }} </li>
                    <li>
                        <a href="javascript: void(0);" class="has-arrow waves-effect">
                            <i class="ri-home-2-line"></i>
                            <span>其他</span>
                        </a>
                    <ul>
                    <li><a href="{{ route('posts.index') }}">文章列表</a></li>
                    <li><a href="{{ route('postCategories.index') }}">文章分類</a></li>
`
var sidemenu_footer = `
                </ul>
                <li class="menu-title">帳號設定 </li>
                    <li>
                        <a href="javascript: void(0);" class="has-arrow waves-effect">
                            <i class="ri-admin-line"></i>
                            <span>{{ __('menu.admin_config') }}</span>
                        </a>
                        <ul class="sub-menu" aria-expanded="false">
                            @can('users-list')
                                <li><a href="{{ route('users.index') }}">{{ __('menu.admin_manage') }}</a></li>
                            @endcan
                            @can('roles-list')
                                <li><a href="{{ route('roles.index') }}">{{ __('menu.role_manage') }}</a></li>
                            @endcan
                        </ul>
                    </li>
                  </li>
            </ul>
        </div>
    </div>
</div>`
