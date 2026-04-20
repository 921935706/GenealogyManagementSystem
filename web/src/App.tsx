import { BrowserRouter as Router, Routes, Route, useLocation, Navigate, useNavigate } from 'react-router-dom'
import { AppShell, Burger, NavLink, Title } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import { IconHome, IconUsers, IconTree, IconPhoto, IconSettings } from '@tabler/icons-react'
import { PersonList, PersonDetail } from '@/features/person'

function Navigation() {
  const [opened, { toggle }] = useDisclosure()
  const location = useLocation()
  const navigate = useNavigate()

  const navItems = [
    { label: '首页', icon: IconHome, path: '/' },
    { label: '人员管理', icon: IconUsers, path: '/persons' },
    { label: '族谱关系', icon: IconTree, path: '/relationships' },
    { label: '媒体资料', icon: IconPhoto, path: '/media' },
    { label: '系统设置', icon: IconSettings, path: '/settings' },
  ]

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: 'sm', collapsed: { mobile: !opened } }}
      padding="md"
    >
      <AppShell.Header>
        <div className="flex items-center h-full px-4">
          <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
          <Title order={3} className="ml-4">族谱管理系统</Title>
        </div>
      </AppShell.Header>

      <AppShell.Navbar p="md">
        {navItems.map((item) => (
          <NavLink 
            key={item.path}
            label={item.label} 
            leftSection={<item.icon size={18} />}
            active={location.pathname === item.path}
            onClick={() => navigate(item.path)}
          />
        ))}
      </AppShell.Navbar>

      <AppShell.Main>
        <Routes>
          <Route path="/" element={
            <div className="p-4">
              <h1 className="text-2xl font-bold mb-4">欢迎使用族谱管理系统</h1>
              <p className="text-gray-600">请选择左侧菜单开始使用系统功能。</p>
            </div>
          } />
          <Route path="/persons" element={<PersonList />} />
          <Route path="/persons/:id" element={<PersonDetail />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </AppShell.Main>
    </AppShell>
  )
}

function App() {
  return (
    <Router>
      <Navigation />
    </Router>
  )
}

export default App